import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrganazerProfileComponent } from './organazer-profile.component';

describe('OrganazerProfileComponent', () => {
  let component: OrganazerProfileComponent;
  let fixture: ComponentFixture<OrganazerProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [OrganazerProfileComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OrganazerProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
